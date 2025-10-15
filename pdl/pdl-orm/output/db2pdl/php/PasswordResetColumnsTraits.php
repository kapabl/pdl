<?php

namespace Com\Mh\Mimanjar\Domain\Data;

trait PasswordResetColumnsTraits
{
    public function createdAt()
    {
        $this->addColumn( 'created_at' );
        return $this;
    }

    public function email()
    {
        $this->addColumn( 'email' );
        return $this;
    }

    public function token()
    {
        $this->addColumn( 'token' );
        return $this;
    }

}

