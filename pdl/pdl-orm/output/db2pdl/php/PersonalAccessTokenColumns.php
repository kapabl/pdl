<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\ColumnsDefinition;

class PersonalAccessTokenColumns extends ColumnsDefinition
{
    protected function setup()
    {
        $this->addColumn( 'abilities', 'string' );
        $this->addColumn( 'created_at', 'string' );
        $this->addColumn( 'id', 'int' );
        $this->addColumn( 'last_used_at', 'string' );
        $this->addColumn( 'name', 'string' );
        $this->addColumn( 'token', 'string' );
        $this->addColumn( 'tokenable_id', 'int' );
        $this->addColumn( 'tokenable_type', 'string' );
        $this->addColumn( 'updated_at', 'string' );
    }
}

