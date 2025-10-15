<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait TeamColumnsTraits
{
    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function id()
    {
        $this->addColumn( 'id' );
        return $this;
    }

    public function name()
    {
        $this->addColumn( 'name' );
        return $this;
    }

    public function personalTeam()
    {
        $this->addColumn( 'personal_team' );
        return $this;
    }

    public function updatedAt()
    {
        $this->addColumn( 'updated_at' );
        return $this;
    }

    public function userId()
    {
        $this->addColumn( 'user_id' );
        return $this;
    }

}

