<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class CourierOrderColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'courier_uuid', 'string' );
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'order_id', 'int' );
        $this->addColumn( 'status', 'string' );
        $this->addColumn( 'updated_at', 'string' );
    }
}

