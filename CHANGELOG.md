# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.4.2] - 2024-10-07

### Changed
- Changed proposal snapshot type from string to integer

## [0.4.1] - 2024-03-19

### Added
- Added method for fetching mci messages

## [0.4.0] - 2024-03-01

### Changed
- Changed the path name of the go module
- Updated go version to 1.21
- Added badges with link to the license and passed workflows

### Added
- Added LICENSE information
- Added info for contributing
- Added github issues templates
- Added linter and unit-tests workflows for github actions

## [0.3.3] - 2024-02-27

### Added
- Added SDK get vote by id method

## [0.3.2] - 2023-12-04

### Added
- Added SDK vote methods

## [0.3.1] - 2023-09-12

### Changed
- Mark votes choice field as json.RawMessage due to multiple values

## [0.3.0] - 2023-08-23

### Added
- Added SDK option for passing the API Key

### Changed
- Changed tests for proposals (move tests to right place)

## [0.2.1] - 2023-07-16

### Added
- Added created_after filter for getting votes

## [0.2.0] - 2023-07-16

### Added
- Added sdk methods for fetching votes
- Added id_in parameter for fetching proposals

### Changed
- Changed structure for request options - it's func now

## [0.1.1] - 2023-07-06

### Added
- Added query for fetching networks
- Added parameter network for fetching ranked space identifier
- Added additional fields for spaces

## [0.1.0] - 2023-07-03

### Added
- Added base graphql queries for Snapshot.org
- Added gqlgenc configuration
- Added SDK for Snapshot.org
- Added linter
