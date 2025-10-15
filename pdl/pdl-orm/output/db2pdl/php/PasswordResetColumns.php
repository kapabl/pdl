<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class PasswordResetColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'email', 'string' );
        $this->addColumn( 'token', 'string' );
    }
}

