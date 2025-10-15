<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\FieldList;

class DiscountInstanceFieldList extends FieldList
{
    protected function setup()
    {
        $this->addField( 'created_at' );
        $this->addField( 'discount_id' );
        $this->addField( 'id' );
        $this->addField( 'instance_info' );
        $this->addField( 'product_id' );
        $this->addField( 'store_id' );
        $this->addField( 'updated_at' );
    }
}

