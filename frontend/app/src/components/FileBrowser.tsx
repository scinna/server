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
function OrderElements(elements: ModelFile[]): ModelFile[] {
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

/**
 * Temporary folder generator for testing purposes
 */
const FakeFolder = (): ModelFile => {
    let root = new ModelFile('/', true);

    let folder1 = new ModelFile('My games screenshots', true, root);
    let subFolder1 = new ModelFile('Minecraft', true, folder1);
    subFolder1.content = [
        new ModelFile('ScreenMC1', false, subFolder1),
        new ModelFile('ScreenMC2', false, subFolder1),
        new ModelFile('ScreenMC3', false, subFolder1),
        new ModelFile('ScreenMC4', false, subFolder1),
    ];

    let subFolder2 = new ModelFile('Gmod', true, folder1);
    subFolder2.content = [
        new ModelFile('ScreenGmod1', false, subFolder2),
        new ModelFile('ScreenGmod2', false, subFolder2),
        new ModelFile('ScreenGmod3', false, subFolder2),
        new ModelFile('ScreenGmod4', false, subFolder2),
    ];

    folder1.content = [
        subFolder1,
        subFolder2,
        new ModelFile('ScreenHLA4', false, folder1),
        new ModelFile('ScreenGTA4', false, folder1),
    ];

    let folder2 = new ModelFile('My web screenshots', true, root);
    let subFolder3 = new ModelFile('Facebook', true, folder2);
    subFolder3.content = [
        new ModelFile('ScreenFB1', false, subFolder3),
        new ModelFile('ScreenFB2', false, subFolder3),
        new ModelFile('ScreenFB3', false, subFolder3),
        new ModelFile('ScreenFB4', false, subFolder3),
    ];

    let subFolder4 = new ModelFile('Twitter', true, folder1);
    subFolder4.content = [
        new ModelFile('ScreenTwitter1', false, subFolder4),
        new ModelFile('ScreenTwitter2', false, subFolder4),
        new ModelFile('ScreenTwitter3', false, subFolder4),
        new ModelFile('ScreenTwitter4', false, subFolder4),
    ];

    folder2.content = [
        subFolder3,
        subFolder4,
        new ModelFile('OxodaoFR', false, folder2),
    ];

    let folder3 = new ModelFile('My work screenshots', true, root);
    folder3.content = [
        new ModelFile('NotMuch_Lol', false, folder3),
    ];

    root.content = [folder1, folder2, folder3];

    return root;
}

const initialState = {
    CurrentFolder: FakeFolder(),
}

export default function() {
    const [state, setState] = React.useState(initialState);
    const classes = useStyles();

    const clickOnIcon = (element: ModelFile|null) => (): any => {
        if (element !== null) {
            if (element?.isFolder)
                setState({ ...state, CurrentFolder: element })
        } else {
            if (state.CurrentFolder !== null && !state.CurrentFolder.IsRoot()) {
                // We already check that the parent is not null
                // @ts-ignore
                setState({ ...state, CurrentFolder: state.CurrentFolder.parent });
            }
        }
    }

    let elements = [];
    if (!state.CurrentFolder.IsRoot())
        elements.push(<File className={classes.icon} key="GoBack" isGoBack={true} onClick={clickOnIcon(null)} />)

    let models = OrderElements(state.CurrentFolder.content);
    models.forEach(elt => {
        elements.push(<File key={elt.filename} className={classes.icon} file={elt} onClick={clickOnIcon(elt)}/>)
    })

    let parents = [];
    if (!state.CurrentFolder.IsRoot())
        parents.push(<Typography color="textPrimary" key="current_folder" className={classes.link}>{state.CurrentFolder.filename}</Typography>)
    
    let currParent: ModelFile|null = state.CurrentFolder.parent;
    while (currParent != null) {
        if (currParent.parent != null) {
            parents.push(<Link color="inherit" key={currParent.filename} href="#" className={classes.link}>{currParent.filename}</Link>)
        }

        currParent = currParent.parent;
    }

    parents.push(<Link color="inherit" href="#" key="root" className={classes.link}>
                    <HomeIcon className={classes.breadcrumbIcon} /> Home
                </Link>)

    parents.reverse();

    return <div className={classes.rootElement}>
        <Breadcrumbs aria-label="breadcrumb">
            {parents}            
        </Breadcrumbs>

        <div className={classes.iconList}>
            {elements}
        </div>
    </div>
}