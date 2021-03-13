import React           from 'react';

import {ProfileEditor} from "../components/ProfileEditor";
import {TokenLister}   from "../components/TokenLister";

import '../assets/scss/Profile.scss';

export function Profile() {
    return <div className="centeredBlock profile">
        <ProfileEditor />
        <TokenLister />
    </div>
}