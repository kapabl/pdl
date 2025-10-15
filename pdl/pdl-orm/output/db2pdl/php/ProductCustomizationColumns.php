<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class ProductCustomizationColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'default_value', 'string' );
        $this->addColumn( 'description', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'is_option', 'string' );
        $this->addColumn( 'is_sold_out', 'string' );
        $this->addColumn( 'max_quantity', 'int' );
        $this->addColumn( 'min_quantity', 'int' );
        $this->addColumn( 'name', 'string' );
        $this->addColumn( 'nutritional_info', 'string' );
        $this->addColumn( 'pickup_price', 'int' );
        $this->addColumn( 'price', 'int' );
        $this->addColumn( 'show_in_order', 'string' );
        $this->addColumn( 'status', 'string' );
        $this->addColumn( 'type', 'string' );
        $this->addColumn( 'updated_at', 'string' );
        $this->addColumn( 'user_id', 'int' );
        $this->addColumn( 'uuid', 'string' );
    }
}

