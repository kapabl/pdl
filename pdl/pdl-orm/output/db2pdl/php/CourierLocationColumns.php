<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class CourierLocationColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'accuracy', 'float' );
        $this->addColumn( 'altitude', 'float' );
        $this->addColumn( 'altitude_accuracy', 'float' );
        $this->addColumn( 'courier_uuid', 'string' );
        $this->addColumn( 'heading', 'float' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'lat', 'float' );
        $this->addColumn( 'location_time', 'string' );
        $this->addColumn( 'lon', 'float' );
        $this->addColumn( 'speed', 'float' );
    }
}

