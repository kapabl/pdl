<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\Where;

class PasswordResetWhere extends Where
{
    protected function setup()
    {
        $this->addField( 'created_at' );
        $this->addField( 'email' );
        $this->addField( 'token' );
    }
}

