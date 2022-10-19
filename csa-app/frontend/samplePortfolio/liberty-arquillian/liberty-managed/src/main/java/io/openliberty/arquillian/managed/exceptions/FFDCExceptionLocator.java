/*
 * Copyright 2018, IBM Corporation, Red Hat Middleware LLC, and individual contributors
 * identified by the Git commit log. 
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package io.openliberty.arquillian.managed.exceptions;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FilenameFilter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.sql.Timestamp;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.logging.Logger;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import io.openliberty.arquillian.managed.exceptions.NestedExceptionBuilder.ExMsg;

/**
 * Finds deployment exceptions by reading FFDC files
 */
public class FFDCExceptionLocator implements DeploymentExceptionLocator {

    private static Logger log = Logger.getLogger(FFDCExceptionLocator.class.getName());

    private final String logsDirectory;
    
    public FFDCExceptionLocator(String logsDirectory) {
        this.logsDirectory = logsDirectory;
    }

    @Override
    public Throwable getException(String appName, String logLine, long deploymentTime) {
        // Get the set of ffdc files
        File[] ffdcFiles = getFfdcFilesSince(deploymentTime);

        // Get StateChangeException FFDC file for this application
        File ffdc = findStateChangeExceptionFfdcFileForApp(appName, ffdcFiles);

        Throwable cause = null;

        if (ffdc != null) {
            // Get StateChangeException ffdc exception names and info chain
            ArrayList<ExMsg> chain = getExceptionCausalChain(ffdc);

            if (chain != null) {
                cause = NestedExceptionBuilder.buildNestedException(chain);
            }
        }

        for (Throwable c = cause; c != null; c = c.getCause()) {
            log.finest("FFDC Exception Chain:" + c.getClass());
        }

        return cause;
    }

    /**
     * Get the set of FFDC files, we are only called when we are expecting to see
     * some FFDC's so we will loop, waiting for 1/100th of a second and up to 10
     * seconds until we see at least on FFDC. Normally we will not have to wait at
     * all.
     * 
     * @return an array of File from the FFDC dir that start "ffdc_"
     * @throws IOException
     */
    private File[] getFfdcFilesSince(Long deployTime) {

        // We want to look for all FFDCs from at least 1 second before deploy until
        // now...
        final long AT_LEAST_ONE_SECOND = (1 * 1000) + 1;
        // Perhaps we have not passed in the deploy time - if so scan all FFDCs:
        final long since = deployTime != null ? (deployTime - AT_LEAST_ONE_SECOND) : 0;
        File[] ffdcFiles = null;

        int attempts = 0;
        do {
            try {
                Thread.sleep(100);
            } catch (InterruptedException e1) {
                // do nothing
            }
            File ffdcDir = null;
            String ffdcDirPath = "NOT_SET";
            ffdcDirPath = logsDirectory + "/ffdc";
            ffdcDir = new File(ffdcDirPath);
            log.finest(
                       "FFDC Dir: " + ffdcDir.getAbsolutePath() + "looking for ffdc_* file since "
                               + new Timestamp(since));
            ffdcFiles = ffdcDir.listFiles(new FilenameFilter() {
                public boolean accept(File dir, String fileName) {
                    File ffdcFile = new File(dir, fileName);
                    return (ffdcFile.lastModified() >= since) && fileName.startsWith("ffdc_");
                }
            });
            attempts++;
            log.finest("FFDC files are: " + Arrays.toString(ffdcFiles));

        } while ((ffdcFiles == null || ffdcFiles.length == 0) && attempts <= 100);

        return ffdcFiles;
    }

    /**
     * Find the app's StateChangeException FFDC file (using Java 6 compatible code).
     * 
     * @param appName
     * @param ffdcFiles
     * 
     * @return The correct FFDC File or null
     * 
     * @throws FileNotFoundException
     * @throws IOException
     */
    private File findStateChangeExceptionFfdcFileForApp(String appName, File[] ffdcFiles) {

        BufferedReader br = null;
        try {
            for (int i = 0; i < ffdcFiles.length; i++) {

                // Set up a bufferedReader
                File ffdcFile = ffdcFiles[i];
                log.finest("Processing FFDC file: " + ffdcFile.getAbsolutePath());
                try {
                    br = new BufferedReader(new InputStreamReader(new FileInputStream(ffdcFile)));
                } catch (FileNotFoundException e) {
                    log.warning("File " + ffdcFile.getAbsolutePath() + " is no longer accessible");
                    continue; // File has been deleted but it might not have been the target one
                }

                // See if it fulfills the criteria
                try {

                    boolean fileIsForStateChangeException = false;
                    boolean fileIsForCorrectApplication = false;

                    String line;
                    while ((line = br.readLine()) != null) {

                        fileIsForStateChangeException = fileIsForStateChangeException || line
                                .contains("Stack Dump = com.ibm.ws.container.service.state.StateChangeException");
                        log.finest("fileIsForStateChangeException:" + fileIsForStateChangeException + " " + line);

                        fileIsForCorrectApplication = fileIsForCorrectApplication
                                || line.contains("appName") && line.contains(appName);
                        log.finest("fileIsForCorrectApplication:" + fileIsForCorrectApplication + "(" + appName + ")"
                                + " " + line);

                        // Have we found all the criteria we are looking for?
                        if (fileIsForStateChangeException && fileIsForCorrectApplication) {
                            // We have found what we are looking for
                            log.finest("Found StateChangeExceptionFfdc for " + appName + ": "
                                    + ffdcFile.getAbsolutePath());
                            return ffdcFile;
                        }

                    }
                    log.finest("Did not find StateChangeExceptionFfdc for " + appName + ": "
                            + ffdcFile.getAbsolutePath() + "(file was scanned completely)");

                } catch (IOException e) {
                    log.warning("File " + ffdcFile.getAbsolutePath() + " is not readable");
                    continue; // File has perhaps deleted but not necessarily a disaster if we find what we
                              // need later
                }
            }

        } finally {
            if (br != null) {
                try {
                    br.close();
                } catch (IOException e) {
                    log.warning(e.getMessage());
                }
            }
        }

        // We did not find what we are looking for
        log.finest("No StateChangeExceptionFfdc for " + appName + " was found ");
        return null;
    }

    /**
     * Get a list of ExMsg objects representing Exceptions in a FFDC log file
     * 
     * @param ffdc
     *            the ffdc file being processed
     * @return an ordered list of exceptions, as ExMsg, foundß
     */
    private ArrayList<ExMsg> getExceptionCausalChain(File ffdc) {

        ArrayList<ExMsg> result = new ArrayList<ExMsg>();

        BufferedReader br = null;
        try {
            br = new BufferedReader(new InputStreamReader(new FileInputStream(ffdc)));
            String line;
            // FFDC files have the form:
            // Stack Dump = com.ibm.ws.container.service.state.StateChangeException:
            // Caused by: org.jboss.weld.exceptions.WeldException:
            // Caused by: org.jboss.weld.exceptions.DefinitionException:

            while ((line = br.readLine()) != null) {
                Pattern p = Pattern.compile("(Stack Dump = |Caused by: )([\\p{L}\\p{N}_$\\.]+?)(:|$)(.*)");

                Matcher m = p.matcher(line);
                if (m.matches()) {
                    ExMsg x = new ExMsg();
                    x.exName = m.group(2);
                    x.logMsg = line;
                    log.finest("match on line: " + line + " resolved to className " + x.exName);
                    result.add(x);
                } else {
                    log.finest("no match found on line: " + line);
                }
            }

        } catch (IOException e) {
            // We are OK to return empty handed (and fail the test if needed) but will log
            // the event...
            log.warning("FFDC file " + ffdc != null ? ffdc.getAbsolutePath()
                    : "null" + " IO Exception: " + e.toString());
        } finally {
            if (br != null) {
                try {
                    br.close();
                } catch (IOException e) {
                    log.warning(e.getMessage());
                }
            }
        }
        log.finest("Exception stack found was: " + Arrays.deepToString(result.toArray()));
        return result;
    }

}
