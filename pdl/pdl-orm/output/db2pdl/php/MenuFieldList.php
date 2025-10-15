<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\FieldList;

class MenuFieldList extends FieldList
{
    protected function setup()
    {
        $this->addField( 'category_ids' );
        $this->addField( 'created_at' );
        $this->addField( 'description' );
        $this->addField( 'id' );
        $this->addField( 'name' );
        $this->addField( 'schedule' );
        $this->addField( 'status' );
        $this->addField( 'store_id' );
        $this->addField( 'updated_at' );
    }
}

