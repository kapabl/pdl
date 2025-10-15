<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class AddressColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'address1', 'string' );
        $this->addColumn( 'address2', 'string' );
        $this->addColumn( 'city', 'string' );
        $this->addColumn( 'country', 'string' );
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'default_delivery', 'string' );
        $this->addColumn( 'default_pickup', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'is_test', 'string' );
        $this->addColumn( 'lat', 'float' );
        $this->addColumn( 'lon', 'float' );
        $this->addColumn( 'name', 'string' );
        $this->addColumn( 'phone', 'string' );
        $this->addColumn( 'state', 'string' );
        $this->addColumn( 'status', 'string' );
        $this->addColumn( 'updated_at', 'string' );
        $this->addColumn( 'user_id', 'int' );
        $this->addColumn( 'zipcode', 'string' );
    }
}

