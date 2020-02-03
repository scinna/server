import React from 'react';
import {Link} from 'react-router-dom';

export default function() {
    return <div className="card above">
        <h4>Welcome to Scinna</h4>
        <div className="content">
            <p>We are going to work together to get you started.</p>
            <p>But first, I will need a few informations to make the server start.</p>
            <p>Here's what we are going to do:</p>
            <ul>
                <li>Set the database up</li>
                <li>Set the mails up</li>
                <li>Configure the server</li>
                <li>Create the admin account</li>
            </ul>
        </div>
        <div className="footer">
            <Link className="btn" to="/database">Let's go</Link>
        </div>
    </div>;
}