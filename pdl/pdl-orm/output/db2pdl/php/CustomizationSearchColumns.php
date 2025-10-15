<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class CustomizationSearchColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'customization_full_name', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'name', 'string' );
        $this->addColumn( 'relation_id', 'int' );
        $this->addColumn( 'updated_at', 'string' );
        $this->addColumn( 'user_id', 'int' );
    }
}

