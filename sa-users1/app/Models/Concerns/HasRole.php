<?php

namespace App\Models\Concerns;

use App\Services\Roles\Role;

trait HasRole
{
    public function getRoleId()
    {
        return $this->role_id;
    }

    public function getRole(): Role
    {
        return Role::from($this->getRoleId());
    }
}
