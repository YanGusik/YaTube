<?php

namespace App\Services\Roles;

enum Role: int
{
    case User = 0;
    case Admin = 1;
}
