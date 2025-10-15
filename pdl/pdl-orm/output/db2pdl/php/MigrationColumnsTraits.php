<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait MigrationColumnsTraits
{
    public function batch()
    {
        $this->addColumn( 'batch' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function migration()
    {
        $this->addColumn( 'migration' );
        return $this;
    }

}

