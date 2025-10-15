<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class ProductOrderColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'position', 'int' );
        $this->addColumn( 'product_id', 'int' );
        $this->addColumn( 'updated_at', 'string' );
        $this->addColumn( 'user_id', 'int' );
    }
}

