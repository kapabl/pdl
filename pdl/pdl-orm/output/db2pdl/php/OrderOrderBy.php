<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\OrderBy;

class OrderOrderBy extends OrderBy
{
    protected function setup()
    {
        $this->addField( 'courier_uuid' );
        $this->addField( 'created_at' );
        $this->addField( 'delivery_tax_amount' );
        $this->addField( 'discount_amount' );
        $this->addField( 'id' );
        $this->addField( 'model' );
        $this->addField( 'notes' );
        $this->addField( 'order_id' );
        $this->addField( 'seller_tax_amount' );
        $this->addField( 'seller_total' );
        $this->addField( 'status' );
        $this->addField( 'store_id' );
        $this->addField( 'sub_total' );
        $this->addField( 'tax_amount' );
        $this->addField( 'total' );
        $this->addField( 'updated_at' );
        $this->addField( 'user_id' );
    }
}

