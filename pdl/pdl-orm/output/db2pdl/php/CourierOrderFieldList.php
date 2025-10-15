<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\FieldList;

class CourierOrderFieldList extends FieldList
{
    protected function setup()
    {
        $this->addField( 'courier_uuid' );
        $this->addField( 'created_at' );
        $this->addField( 'id' );
        $this->addField( 'order_id' );
        $this->addField( 'status' );
        $this->addField( 'updated_at' );
    }
}

