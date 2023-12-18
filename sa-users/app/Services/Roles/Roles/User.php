<?php

namespace App\Services\Roles\Roles;

use App\Services\Roles\Permissions\VideoPermission;
use App\Services\Roles\Role;

class User extends BaseRole
{
    public function getRole(): Role
    {
        return Role::User;
    }

    public function getPermissions(): array
    {
        return [
            VideoPermission::CAN_CREATE_VIDEO,
            VideoPermission::CAN_UPDATE_VIDEO,
            VideoPermission::CAN_DELETE_VIDEO,
        ];
    }
}
