<?php

namespace App\Services\Roles\Roles;

use App\Services\Roles\Permissions\VideoPermission;
use App\Services\Roles\Role;

class Admin extends User
{
    public function getRole(): Role
    {
        return Role::Admin;
    }

    public function getPermissions(): array
    {
        return array_merge(parent::getPermissions(), [
            VideoPermission::CAN_ADMIN_CREATE_VIDEO,
            VideoPermission::CAN_ADMIN_UPDATE_VIDEO,
            VideoPermission::CAN_ADMIN_DELETE_VIDEO,
        ]);
    }
}
