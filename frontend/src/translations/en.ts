export default {
    visibility: {
        public: 'Public',
        unlisted: 'Unlisted',
        private: 'Private'
    },
    menu: {
        home: 'Home',
        login: 'Login',
        register: 'Register',
        server: 'Server',
        logout: 'Logout'
    },
    errors: {
        unknown: 'Unknown error',
        passwordNotMatching: 'Password do not match',
    },
    registration: {
        title: "Register",
        username: "Username",
        email: "Email",
        password: "New password",
        repeat_password: "Repeat password",
        invite_code: "Invite code",
        button: "Register",
    },
    login: {
        title: 'Login',
        username: 'Username',
        password: 'Password',
        button: 'Login'
    },
    validate: {
        accountValidated: 'Your account is now validated',
        connect: 'You can now login with the username'
    },
    my_profile: {
        account: {
            tab_name: 'My account',
            current_password: 'Current password',
            new_password: 'New password',
            update: 'Save',
            success: 'Profile has been successfully updated',
        },
        tokens: {
            tab_name: 'Connected devices',
            logged_at: 'Logged in at ',
            revoked_at: 'Revoked at',
            last_seen: 'Last seen',
            never: 'Never',
            on: 'on',
            revoke_dialog: {
                title: 'Revoke a token',
                text: 'Revoking this token will disconnect you from all devices / app using it.',
                revoke: 'Revoke',
                cancel: 'Cancel'
            }
        },
        sharex: {
            tab_name: 'ShareX'
        }
    },
    server_settings: {
        invite: {
            tab_name: "Invite code",
            code: "Invite code",
            generate: "Generate",
            list: {
                generated_by: 'Generated by',
                on: 'On',
            },
            delete_dialog: {
                title: 'Remove an invite code',
                text: 'Removing this invite code will prevent anyone from using it to register.',
                remove: 'Remove',
                cancel: 'Cancel'
            }
        }
    },
    dial: {
      textbin: 'Textbin',
      url_shortener: 'URL shortener'
    },
    browser: {
        create_folder: {
            title: 'Create a folder',
            folder_name: 'Name',
            visibility: 'Visibility',
            create: 'Create',
            cancel: 'Cancel'
        },
        file_uploader: {
            title: 'Upload',
            file_name: 'Name',
            description: 'Description',
            visibility: 'Visibility',
            cancel: 'Cancel',
            upload: 'Upload'
        },
        context: {
            edit: 'Edit',
            remove: 'Delete',
        },
        modals: {
            remove_collection: {
                title: 'Deleting collection',
                text: 'Are you sure you want to remove this collection? It will remove any files in it.',
                cancel: 'Cancel',
                delete: 'Delete'
            },
            remove_media: {
                title: 'Deleting media',
                text: 'Are you sure you want to delete this media?',
                cancel: 'Cancel',
                delete: 'Delete'
            },
            edit_media: {
                title: 'Editing media',
                media_title: 'Title',
                description: 'Description',
                visibility: 'Visibility'
            },
            remove_link: {
                title: 'Removing link',
                text: 'This will remove your link pointing to '
            }
        }
    },
    shortener: {
        link: 'Link',
        send: 'Minify!',
        scinna_link: 'Scinna link',
        amt_views: 'Views'
    }
};