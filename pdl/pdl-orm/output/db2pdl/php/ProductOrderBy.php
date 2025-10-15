<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\OrderBy;

class ProductOrderBy extends OrderBy
{
    protected function setup()
    {
        $this->addField( 'created_at' );
        $this->addField( 'delivery_price' );
        $this->addField( 'deposit_options' );
        $this->addField( 'description' );
        $this->addField( 'details' );
        $this->addField( 'featured' );
        $this->addField( 'id' );
        $this->addField( 'images' );
        $this->addField( 'keywords' );
        $this->addField( 'name' );
        $this->addField( 'price' );
        $this->addField( 'quantity' );
        $this->addField( 'slug' );
        $this->addField( 'status' );
        $this->addField( 'store_id' );
        $this->addField( 'updated_at' );
        $this->addField( 'user_id' );
        $this->addField( 'uuid' );
    }
}

