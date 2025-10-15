<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class CategoryColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'name', 'string' );
        $this->addColumn( 'position', 'int' );
        $this->addColumn( 'slug', 'string' );
        $this->addColumn( 'status', 'string' );
        $this->addColumn( 'store_id', 'int' );
        $this->addColumn( 'updated_at', 'string' );
    }
}

