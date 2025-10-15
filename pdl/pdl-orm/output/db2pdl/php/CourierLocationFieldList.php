<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\FieldList;

class CourierLocationFieldList extends FieldList
{
    protected function setup()
    {
        $this->addField( 'accuracy' );
        $this->addField( 'altitude' );
        $this->addField( 'altitude_accuracy' );
        $this->addField( 'courier_uuid' );
        $this->addField( 'heading' );
        $this->addField( 'id' );
        $this->addField( 'lat' );
        $this->addField( 'location_time' );
        $this->addField( 'lon' );
        $this->addField( 'speed' );
    }
}

