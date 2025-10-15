<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class CacheColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'expiration', 'int' );
        $this->addColumn( 'key', 'string' );
        $this->addColumn( 'value', 'string' );
    }
}

