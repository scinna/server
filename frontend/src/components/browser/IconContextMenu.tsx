import {ListItemIcon, Menu, MenuItem, Typography} from "@material-ui/core";
import {Delete as DeleteIcon, Edit as EditIcon}   from "@material-ui/icons";
import React                                      from "react";
import {useIconContextMenu}                       from "../../context/IconContextMenuProvider";
import {useModal}                                 from "../../context/ModalProvider";
import {DeleteCollection}                         from "../modals/DeleteCollection";
import i18n                                       from "i18n-js";
import {DeleteMedia}                              from "../modals/DeleteMedia";
import {FolderEditor}                             from "../modals/FolderEditor";


export function IconContextMenu() {
    const {
        isVisible,
        mouseX,
        mouseY,
        hide,
        collection,
        media
    } = useIconContextMenu();
    const modal = useModal();

    if (!isVisible) {
        return null;
    }

    return <Menu
        keepMounted
        open={mouseY !== null}
        onClose={hide}
        anchorReference="anchorPosition"
        anchorPosition={{top: mouseY ?? 0, left: mouseX ?? 0}}
    >
        <MenuItem onClick={() => modal.show(
            collection !== null
                ? <FolderEditor collection={collection} closeCallback={hide} />
                : media !== null
                    ? null
                    : null
        )}>
            <ListItemIcon>
                <EditIcon fontSize="small"/>
            </ListItemIcon>
            <Typography variant="inherit">{i18n.t('browser.context.edit')}</Typography>
        </MenuItem>
        <MenuItem onClick={() => modal.show(
            collection !== null ?
                <DeleteCollection collection={collection} successCallback={hide}/>
                : media !== null ?
                    <DeleteMedia media={media} successCallback={hide}/>
                    : null
            )}>
            <ListItemIcon>
                <DeleteIcon fontSize="small" color="secondary"/>
            </ListItemIcon>
            <Typography variant="inherit" color="secondary">{i18n.t('browser.context.remove')}</Typography>
        </MenuItem>
    </Menu>
}