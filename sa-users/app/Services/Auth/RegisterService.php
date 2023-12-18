<?php

namespace App\Services\Auth;

use App\Models\User;
use App\Services\Auth\TDU\UserRegisterTDU;
use Illuminate\Support\Facades\Hash;

class RegisterService
{
    public function createUser(UserRegisterTDU $userTDU): User|false
    {
        if (User::whereEmail($userTDU->email)->exists())
        {
            return false;
        }

        return User::create([
            'email' => $userTDU->email,
            'name' => $userTDU->name,
            'password' => Hash::make($userTDU->password)
        ]);
    }
}
