<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\OrderBy;

class ShoppingcartOrderBy extends OrderBy
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

