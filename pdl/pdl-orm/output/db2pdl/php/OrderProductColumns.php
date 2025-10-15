<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class OrderProductColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'order_id', 'int' );
        $this->addColumn( 'product_id', 'int' );
        $this->addColumn( 'quantity', 'int' );
        $this->addColumn( 'updated_at', 'string' );
    }
}

