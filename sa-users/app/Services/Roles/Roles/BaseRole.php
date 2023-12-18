<?php

namespace App\Services\Roles\Roles;

use App\Services\Roles\Role;

abstract class BaseRole
{
    public function getName(): string
    {
        return $this->getRole()->name;
    }
    public abstract function getRole(): Role;
    public abstract function getPermissions(): array;

    public function hasPermission(string $permission): bool
    {
        $permissions = $this->getPermissions();
        return in_array($permission, $permissions);
    }
}
