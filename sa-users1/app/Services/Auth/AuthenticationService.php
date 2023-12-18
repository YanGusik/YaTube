<?php

namespace App\Services\Auth;

use App\Models\User;
use App\Services\Auth\Exceptions\InvalidClient;
use App\Services\Auth\Exceptions\UserNotFound;
use GuzzleHttp\Client;
use GuzzleHttp\Exception\ClientException;
use GuzzleHttp\Exception\GuzzleException;
use Laravel\Passport\Client as OClient;

class AuthenticationService
{
    /**
     * @throws UserNotFound
     * @throws InvalidClient
     * @throws GuzzleException
     */
    public function signIn(string $email, string $password): array
    {
        return $this->getTokenAndRefreshToken($email, $password);
    }

    // TODO: http://localhost:8080/oauth/token 366 ms (email,password,client secret, client id) Too much to get a token
    // TODO: http://localhost:8080/api/login (custom method) 457 ms (email, password) guzzle reduces performance by 60-90ms, thanks laravel passport
    /**
     * @throws UserNotFound
     * @throws InvalidClient
     * @throws GuzzleException
     */
    private function getTokenAndRefreshToken($email, $password): array
    {
        $passwordClient = OClient::where('password_client', 1)->first();
        if ($passwordClient === null) {
            throw new InvalidClient('Password client not found');
        }

        $http = new Client();
        try {
            $response = $http->request('POST', 'sa_users_webserver/oauth/token', [
                'form_params' => [
                    'grant_type' => 'password',
                    'client_id' => $passwordClient->id,
                    'client_secret' => $passwordClient->secret,
                    'username' => $email,
                    'password' => $password,
                    'scope' => '*',
                ],
            ]);
            return json_decode((string)$response->getBody(), true);
        } catch (ClientException $exception) {
            $jsonBody = json_decode((string)$exception->getResponse()->getBody(), true);
            throw match ($jsonBody['error'] ?? null) {
                'invalid_grant' => new UserNotFound(),
                'invalid_client' => new InvalidClient($jsonBody['error']),
                default => new \Exception($jsonBody['error'] ?? sprintf("status:%s", $exception->getCode())),
            };
        }
    }
}
