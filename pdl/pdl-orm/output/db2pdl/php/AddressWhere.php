<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\Where;

class AddressWhere extends Where
{
    protected function setup()
    {
        $this->addField( 'address1' );
        $this->addField( 'address2' );
        $this->addField( 'city' );
        $this->addField( 'country' );
        $this->addField( 'created_at' );
        $this->addField( 'default_delivery' );
        $this->addField( 'default_pickup' );
        $this->addField( 'id' );
        $this->addField( 'is_test' );
        $this->addField( 'lat' );
        $this->addField( 'lon' );
        $this->addField( 'name' );
        $this->addField( 'phone' );
        $this->addField( 'state' );
        $this->addField( 'status' );
        $this->addField( 'updated_at' );
        $this->addField( 'user_id' );
        $this->addField( 'zipcode' );
    }
}

