<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class ProductCustomizationRelationColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'parent_id', 'int' );
        $this->addColumn( 'position', 'int' );
        $this->addColumn( 'product_customization_id', 'int' );
        $this->addColumn( 'product_id', 'int' );
        $this->addColumn( 'updated_at', 'string' );
        $this->addColumn( 'user_id', 'int' );
    }
}

