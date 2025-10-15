<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\Where;

class CategoryProductWhere extends Where
{
    protected function setup()
    {
        $this->addField( 'category_id' );
        $this->addField( 'created_at' );
        $this->addField( 'id' );
        $this->addField( 'position' );
        $this->addField( 'product_id' );
        $this->addField( 'store_id' );
        $this->addField( 'updated_at' );
    }
}

