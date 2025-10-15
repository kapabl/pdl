<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class MigrationColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'batch', 'int' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'migration', 'string' );
    }
}

