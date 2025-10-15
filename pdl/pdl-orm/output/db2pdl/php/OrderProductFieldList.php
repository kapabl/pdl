<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\FieldList;

class OrderProductFieldList extends FieldList
{
    protected function setup()
    {
        $this->addField( 'created_at' );
        $this->addField( 'id' );
        $this->addField( 'order_id' );
        $this->addField( 'product_id' );
        $this->addField( 'quantity' );
        $this->addField( 'updated_at' );
    }
}

