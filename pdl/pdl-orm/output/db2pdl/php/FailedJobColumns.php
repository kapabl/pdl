<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class FailedJobColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'connection', 'string' );
        $this->addColumn( 'exception', 'string' );
        $this->addColumn( 'failed_at', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'payload', 'string' );
        $this->addColumn( 'queue', 'string' );
        $this->addColumn( 'uuid', 'string' );
    }
}

