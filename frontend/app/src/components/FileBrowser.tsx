import React from 'react';

import File from './File';
import ModelFile from '../model/File';

import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import Breadcrumbs from '@material-ui/core/Breadcrumbs';
import Link from '@material-ui/core/Link';
import HomeIcon from '@material-ui/icons/Home';

const useStyles = makeStyles(theme => ({
    rootElement: {
        width: '100%',
        backgroundColor: '#262628',
        padding: '2em',
        borderRadius: '1em',
        boxShadow: 'inset 0 2px 4px 0 rgba(0, 0, 0, 0.06)',
    },

    link: {
        display: 'flex',
    },

    breadcrumbIcon: {
        marginRight: theme.spacing(0.5),
        width: 20,
        height: 20,
    },

    iconList: {
        display: 'flex',
        flexDirection: 'row',
        flexWrap: 'wrap',
        justifyContent: 'center',   
        margin: '1em',
    },

    icon: {
        marginTop: '1em',
    }
}))

// This methods re-order elements to have folder first, file then, both ordered alphabetically
// Maybe there should be options in the future to order by upload date or other things like that
function OrderElements(elements: ModelFile[]) {
    let ordered: ModelFile[] = [];

    // We sort everything by filename
    ordered.sort(function(a, b) {
        return (a.filename.toLowerCase() < b.filename.toLowerCase()) ? -1 : (a.filename.toLowerCase() > b.filename.toLowerCase()) ? 1 : 0;
    });

    elements.forEach(elt => {
        if (elt.isFolder) {
            ordered.push(elt)
        }
    });

    elements.forEach(elt => {
        if (!elt.isFolder) {
            ordered.push(elt)
        }
    })

    return ordered;
}

export default function() {
    const classes = useStyles();

    let elements = [];
    elements.push(<File className={classes.icon} key="GoBack" isGoBack={true} />)

    let tmpFiles = []
    for (let i = 0; i < 15; i++) {
        let isFolder = i%2===0 && i < 5;
        let txt = (isFolder ? "Folder " : "Item ");
        tmpFiles.push(new ModelFile(txt + i, isFolder))
    }

    tmpFiles = OrderElements(tmpFiles);

    tmpFiles.forEach(elt => {
        elements.push(<File className={classes.icon} file={elt} />)
    })

    return <div className={classes.rootElement}>
        <Breadcrumbs aria-label="breadcrumb">
            <Link color="inherit" href="/" className={classes.link}>
                <HomeIcon className={classes.breadcrumbIcon} />
                Home
            </Link>
            <Link color="inherit" href="/getting-started/installation/" className={classes.link}>MyFolder</Link>
            <Typography color="textPrimary" className={classes.link}>MySubFolder</Typography>
        </Breadcrumbs>

        <div className={classes.iconList}>
            {elements}
        </div>
    </div>
}