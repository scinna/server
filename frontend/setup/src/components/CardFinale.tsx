import React from 'react';

export default function() {
    return <div className="card above">
        <h4>It's ready!</h4>
        <div className="content">
            <p>Congratulation, you've just set your server up and it's ready to be used.</p>
            <p>You may want to install the Scinnapse or the Scinnamon clients now!</p>
            <p>Please restart the server now to run in standard mode. If you have your service auto-restart on stop, this will be done automatically</p>
        </div>
        <div className="footer">
            <a className="btn" href="/">Dive into Scinna!</a>
        </div>
    </div>;
}