<?php

namespace App\Http\Controllers\Auth;

use App\Http\Controllers\Controller;
use App\Http\Requests\Auth\SignInUserRequest;
use App\Models\User;
use App\Services\Auth\AuthenticationService;
use App\Services\Auth\Exceptions\UserNotFound;
use App\Services\Roles\RoleService;
use Illuminate\Http\Request;

class LoginController extends Controller
{
    public function me(Request $request, RoleService $roleService)
    {
        /* @var User $user */
        $user = $request->user();
        $data = [
            'role' => $user->getRole()->name,
            'permissions' => $roleService->getRole($user->getRole())->getPermissions()
        ];
        return array_merge($user->toArray(),$data);
    }

    public function signIn(SignInUserRequest $request, AuthenticationService $service): \Illuminate\Http\JsonResponse
    {
        try {
            $token = $service->signIn($request->validated('email'), $request->validated('password'));
            return response()->json(['data' => $token]);
        }
        catch (UserNotFound $exception)
        {
            return response()->json(['message' => 'Unauthorized'], 401);
        }
    }
}
