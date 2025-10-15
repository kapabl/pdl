<?php

namespace Com\Mh\Mimanjar\Domain\Data;

use Com\Mh\Ds\Infrastructure\Data\Db\Where;

class PersonalAccessTokenWhere extends Where
{
    protected function setup()
    {
        $this->addField( 'abilities' );
        $this->addField( 'created_at' );
        $this->addField( 'id' );
        $this->addField( 'last_used_at' );
        $this->addField( 'name' );
        $this->addField( 'token' );
        $this->addField( 'tokenable_id' );
        $this->addField( 'tokenable_type' );
        $this->addField( 'updated_at' );
    }
}

