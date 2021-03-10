import React from 'react';
import { useParams } from 'react-router-dom';

interface BrowserParams {
    username: string;
    path: string;
}

export function Browser() {
    const { username, path } = useParams<BrowserParams>();
    return <div>Browsing user { username } at { path }</div>;
}