<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\Where;

class ProductOrderWhere extends Where
{
    protected function setup()
    {
        $this->addField( 'created_at' );
        $this->addField( 'id' );
        $this->addField( 'position' );
        $this->addField( 'product_id' );
        $this->addField( 'updated_at' );
        $this->addField( 'user_id' );
    }
}

