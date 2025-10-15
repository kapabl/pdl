<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class DiscountInstanceColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'discount_id', 'int' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'instance_info', 'string' );
        $this->addColumn( 'product_id', 'int' );
        $this->addColumn( 'store_id', 'int' );
        $this->addColumn( 'updated_at', 'string' );
    }
}

