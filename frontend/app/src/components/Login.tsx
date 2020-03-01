import React from 'react';

import {Grid, Typography, makeStyles } from '@material-ui/core';

const useStyles = makeStyles(theme => ({
    container: {
        display: 'block',
    },
    text: {
        marginTop: '2em',
        display: 'block',
    }
}))

export default function() {
    const classes = useStyles();

    return <Grid container spacing={6}>
        <Grid container item xs={12} sm={6} className={classes.container}>
            <Typography color="textPrimary" variant="h3" component="h1">
                Scinna
            </Typography>
            <Typography color="textSecondary" component="p" className={classes.text}>
                This is a Scinna instance.
            </Typography>
            <Typography color="textSecondary" component="p" className={classes.text}>
                Scinna is an open-source, self-hosted screenshot and picture server. Find more about it at <a className="pretty-link" href="https://scinna.app/">https://scinna.app</a>
            </Typography>
        </Grid>
        
        <Grid container item xs={12} sm={6}>
            Scinna login screen
        </Grid>
    </Grid>;
}