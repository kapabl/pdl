<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class ShoppingcartColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'content', 'string' );
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'identifier', 'string' );
        $this->addColumn( 'instance', 'string' );
        $this->addColumn( 'updated_at', 'string' );
    }
}

