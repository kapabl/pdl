<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class SessionColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'id', 'string' );
        $this->addColumn( 'ip_address', 'string' );
        $this->addColumn( 'last_activity', 'int' );
        $this->addColumn( 'payload', 'string' );
        $this->addColumn( 'user_agent', 'string' );
        $this->addColumn( 'user_id', 'int' );
    }
}

