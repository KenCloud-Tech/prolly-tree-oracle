// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

    struct Permission {
        // // Whether the user is allowed to insert.
        // bool allowInsert;
        // // Whether the user is allowed to update.
        // bool allowUpdate;

        // Whether the user is allowed to write.
        bool allowWrite;
        // Whether the user is allowed to query.
        bool allowQuery;
        // Whether the user is allowed to delete.
        bool allowDelete;
    }
