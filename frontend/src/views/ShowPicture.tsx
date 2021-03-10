import React from 'react';
import { useParams } from 'react-router-dom';

interface ShowPictureParams {
    pictureId: string;
}

export function ShowPicture() {
    const { pictureId } = useParams<ShowPictureParams>();

    return <div>Showing picture "{pictureId}"</div>
}