<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class ProductColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'delivery_price', 'int' );
        $this->addColumn( 'deposit_options', 'string' );
        $this->addColumn( 'description', 'string' );
        $this->addColumn( 'details', 'string' );
        $this->addColumn( 'featured', 'int' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'images', 'string' );
        $this->addColumn( 'keywords', 'string' );
        $this->addColumn( 'name', 'string' );
        $this->addColumn( 'price', 'int' );
        $this->addColumn( 'quantity', 'int' );
        $this->addColumn( 'slug', 'string' );
        $this->addColumn( 'status', 'string' );
        $this->addColumn( 'store_id', 'int' );
        $this->addColumn( 'updated_at', 'string' );
        $this->addColumn( 'user_id', 'int' );
        $this->addColumn( 'uuid', 'string' );
    }
}

