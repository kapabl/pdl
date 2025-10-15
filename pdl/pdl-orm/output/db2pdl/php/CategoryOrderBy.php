<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\OrderBy;

class CategoryOrderBy extends OrderBy
{
    protected function setup()
    {
        $this->addField( 'created_at' );
        $this->addField( 'id' );
        $this->addField( 'name' );
        $this->addField( 'position' );
        $this->addField( 'slug' );
        $this->addField( 'status' );
        $this->addField( 'store_id' );
        $this->addField( 'updated_at' );
    }
}

