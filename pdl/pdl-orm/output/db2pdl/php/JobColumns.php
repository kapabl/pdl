<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class JobColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'attempts', 'int' );
        $this->addColumn( 'available_at', 'int' );
        $this->addColumn( 'created_at', 'int' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'payload', 'string' );
        $this->addColumn( 'queue', 'string' );
        $this->addColumn( 'reserved_at', 'int' );
    }
}

