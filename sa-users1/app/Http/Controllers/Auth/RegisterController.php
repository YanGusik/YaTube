<?php

namespace App\Http\Controllers\Auth;

use App\Http\Controllers\Controller;
use App\Http\Requests\Auth\RegisterUserRequest;
use App\Services\Auth\AuthenticationService;
use App\Services\Auth\RegisterService;
use App\Services\Auth\TDU\UserRegisterTDU;
use Illuminate\Support\Facades\DB;

class RegisterController extends Controller
{
    public function register(RegisterUserRequest $request, RegisterService $registerService, AuthenticationService $authenticationService): \Illuminate\Http\JsonResponse
    {
        DB::beginTransaction();
        try {
            $validated = $request->validated();
            $tdu = new UserRegisterTDU($validated['email'], $validated['name'], $validated['password']);
            $user = $registerService->createUser($tdu);
            if ($user === false) {
                return response()->json(['message' => 'User cannot created'], 400);
            }
            $token = $authenticationService->signIn($tdu->email, $tdu->password);
            DB::commit();
            return response()->json(['message' => 'Success', 'data' => $user->toArray(), 'token' => $authenticationService]);
        }
        catch (\Exception $exception)
        {
            DB::rollBack();
            throw $exception;
        }
    }
}
