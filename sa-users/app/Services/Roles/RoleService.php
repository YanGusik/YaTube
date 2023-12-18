<?php

namespace App\Services\Roles;

use App\Services\Roles\Roles\Admin;
use App\Services\Roles\Roles\BaseRole;
use App\Services\Roles\Roles\User;

class RoleService
{
    public const ROLES = [
        Role::User->value => User::class,
        Role::Admin->value => Admin::class,
    ];

    public function getRole(Role $role)
    {
        return new (self::ROLES[$role->value]);
    }
}
