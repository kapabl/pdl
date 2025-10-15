<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class OrderColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'courier_uuid', 'string' );
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'delivery_tax_amount', 'int' );
        $this->addColumn( 'discount_amount', 'int' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'model', 'string' );
        $this->addColumn( 'notes', 'string' );
        $this->addColumn( 'order_id', 'string' );
        $this->addColumn( 'seller_tax_amount', 'int' );
        $this->addColumn( 'seller_total', 'int' );
        $this->addColumn( 'status', 'string' );
        $this->addColumn( 'store_id', 'int' );
        $this->addColumn( 'sub_total', 'int' );
        $this->addColumn( 'tax_amount', 'int' );
        $this->addColumn( 'total', 'int' );
        $this->addColumn( 'updated_at', 'string' );
        $this->addColumn( 'user_id', 'int' );
    }
}

