<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\Where;

class ShoppingcartWhere extends Where
{
    protected function setup()
    {
        $this->addField( 'content' );
        $this->addField( 'created_at' );
        $this->addField( 'identifier' );
        $this->addField( 'instance' );
        $this->addField( 'updated_at' );
    }
}

